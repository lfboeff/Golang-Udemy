package repositorios

import (
	"api_mod/src/modelos"
	"database/sql"
)

// PublicacoesRep representa um repositório de publicações
type PublicacoesRep struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um novo repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *PublicacoesRep {
	return &PublicacoesRep{db}
}

// Criar insere uma publicação no banco de dados
func (repositorioPublicacoes *PublicacoesRep) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, err := repositorioPublicacoes.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	publicacaoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(publicacaoID), nil
}

// Buscar traz do banco de dados as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorioPublicacoes *PublicacoesRep) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {

	rows, err := repositorioPublicacoes.db.Query(`
		select distinct p.*, u.nick
		from publicacoes p
		inner join seguidores s on p.autor_id = s.usuario_id
		inner join usuarios u on u.id = p.autor_id
		where u.id = ? or s.seguidor_id = ?
		order by p.id desc`,
		usuarioID, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publicacoes []modelos.Publicacao

	for rows.Next() {

		var publicacao modelos.Publicacao

		if err = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorioPublicacoes *PublicacoesRep) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {

	row, err := repositorioPublicacoes.db.Query(`
		select p.*, u.nick
		from publicacoes p
		inner join usuarios u
		on u.id = p.autor_id
		where p.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer row.Close()

	var publicacao modelos.Publicacao

	if row.Next() {
		if err = row.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}

	return publicacao, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorioPublicacoes *PublicacoesRep) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {

	statement, err := repositorioPublicacoes.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil
}

// Deletar remove uma publicação do banco de dados
func (repositorioPublicacoes *PublicacoesRep) Deletar(publicacaoID uint64) error {

	statement, err := repositorioPublicacoes.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

// BuscarPorUsuario traz do banco de dados todas as publicações de um usuário específico
func (repositorioPublicacoes *PublicacoesRep) BuscarPorUsuario(usuario_id uint64) ([]modelos.Publicacao, error) {

	rows, err := repositorioPublicacoes.db.Query(`
		select p.*, u.nick
		from publicacoes p
		inner join usuarios u
		on u.id = p.autor_id
		where p.autor_id = ?`,
		usuario_id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publicacoes []modelos.Publicacao

	for rows.Next() {

		var publicacao modelos.Publicacao

		if err = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona no banco de dados uma curtida na publicação
func (repositorioPublicacoes *PublicacoesRep) Curtir(publicacaoID uint64) error {

	statement, err := repositorioPublicacoes.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

// Descurtir subtrai no banco de dados uma curtida da publicação, caso o número de curtidas seja maior que 0
func (repositorioPublicacoes *PublicacoesRep) Descurtir(publicacaoID uint64) error {

	statement, err := repositorioPublicacoes.db.Prepare(`
		update publicacoes set curtidas =
		case
			when curtidas > 0 then curtidas - 1
			else 0
		end
		where id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
