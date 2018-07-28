Feature: Login
 Para que eupossa ter scesso as mnhas tarefas
 Sendo um usuario
 Posso me authenticar com os meus dados previamente cadastrados

  Background: Formulario de login
    Given que eu acessei a pagina pricipal

# To run the @success only here's the run command in terminal
# godog -t @success
  @success
  Scenario: Login de usuario
    When faco login com "eu@papito.io" e "123456"
    Then sou autenticado com sucesso

  Scenario: Senha incorreta
    When faco login com "eu@papito,io" e "xpto123"
    Then devo ver a seguite mensagem "Senha invalida."

  Scenario: Email invalido
    When faco login com "eu@papito.io" e "xpto123"
    Then demo ver a seguite mensagem "Email incorreto ou ausente."
