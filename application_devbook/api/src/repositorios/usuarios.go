package repositorios

import (
	"api_mod/src/modelos"
	"database/sql"
	"fmt"
)

// UsuariosRep representa um repositório de usuários
type UsuariosRep struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um novo repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *UsuariosRep {
	return &UsuariosRep{db}
}

// Criar insere um usuário no banco de dados e retorna seu ID (uint64) e erro (error)
func (repositorioUsuarios UsuariosRep) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, err := repositorioUsuarios.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	usuarioId, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(usuarioId), nil

}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorioUsuarios UsuariosRep) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // "%nomeOunick%"

	rows, err := repositorioUsuarios.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []modelos.Usuario

	for rows.Next() {
		var user modelos.Usuario

		if err = rows.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, user)
	}

	return usuarios, nil

}

// BuscarPorID traz um usuário do banco de dados
func (repositorioUsuarios UsuariosRep) BuscarPorID(usuarioID uint64) (modelos.Usuario, error) {

	rows, err := repositorioUsuarios.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", usuarioID)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer rows.Close()

	var usuario modelos.Usuario

	if rows.Next() {
		if err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil

}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorioUsuarios UsuariosRep) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {

	statement, err := repositorioUsuarios.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); err != nil {
		return err
	}

	return nil

}

// Deletar remove as informações de um usuário do banco de dados
func (repositorioUsuarios UsuariosRep) Deletar(usuarioID uint64) error {

	statement, err := repositorioUsuarios.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID); err != nil {
		return err
	}

	return nil

}

// BuscarPorEmail busca um usuário por email e retorna o seu ID e senha com hash para fazer o login
func (repositorioUsuarios UsuariosRep) BuscarPorEmail(email string) (modelos.Usuario, error) {

	rows, err := repositorioUsuarios.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer rows.Close()

	var usuario modelos.Usuario

	if rows.Next() {
		if err = rows.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil

}

// Seguir permite que um usuário siga um outro
func (repositorioUsuarios UsuariosRep) Seguir(usuarioASeguirID, seguidorID uint64) error {

	statement, err := repositorioUsuarios.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioASeguirID, seguidorID); err != nil {
		return err
	}

	return nil
}

// PararDeSeguir permite que um usuário pare de seguir um outro
func (repositorioUsuarios UsuariosRep) PararDeSeguir(usuarioPararSeguirID, seguidorID uint64) error {

	statement, err := repositorioUsuarios.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioPararSeguirID, seguidorID); err != nil {
		return err
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorioUsuarios UsuariosRep) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	rows, err := repositorioUsuarios.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u
	inner join seguidores s on u.id = s.seguidor_id
	where s.usuario_id = ?`, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seguidores []modelos.Usuario

	for rows.Next() {

		var seguidor modelos.Usuario

		if err = rows.Scan(&seguidor.ID, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.CriadoEm); err != nil {
			return nil, err
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorioUsuarios UsuariosRep) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	rows, err := repositorioUsuarios.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm
	from usuarios u
	inner join seguidores s on u.id = s.usuario_id
	where s.seguidor_id = ?`, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuariosSeguidos []modelos.Usuario

	for rows.Next() {

		var usuarioSeguido modelos.Usuario

		if err = rows.Scan(&usuarioSeguido.ID, &usuarioSeguido.Nome, &usuarioSeguido.Nick, &usuarioSeguido.Email, &usuarioSeguido.CriadoEm); err != nil {
			return nil, err
		}

		usuariosSeguidos = append(usuariosSeguidos, usuarioSeguido)
	}

	return usuariosSeguidos, nil
}

// BuscarSenha traz a senha de um usuário a partir do seu ID
func (repositorioUsuarios UsuariosRep) BuscarSenha(usuarioID uint64) (string, error) {

	row, err := repositorioUsuarios.db.Query("select senha from usuarios where id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	// var usuario modelos.Usuario
	senha := ""

	if row.Next() {
		// if err = row.Scan(&usuario.Senha); err != nil {
		if err = row.Scan(&senha); err != nil {
			return "", nil
		}
	}

	// return usuario.Senha, nil
	return senha, nil
}

// AtualizarSenha registra a nova senha do usuário
func (repositorioUsuarios UsuariosRep) AtualizarSenha(usuarioID uint64, senhaNovaComHash string) error {

	statement, err := repositorioUsuarios.db.Prepare("update usuarios set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senhaNovaComHash, usuarioID); err != nil {
		return err
	}

	return nil
}
