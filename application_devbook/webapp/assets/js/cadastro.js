$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função criarUsuario")

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("Passwords do not match!");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nome: $('#email').val(),
            nome: $('#nick').val(),
            nome: $('#senha').val()
        }
    });
}