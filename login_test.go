package login

import (
	"fmt"
	"regexp"
	"strings"

	"login/support"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func queEnAcesseiAPaginPrincipal() error {
	Driver.Get("https://marktests.herokuapp.com")
	return nil
}

func faoLoginComE(email, senha string) (err error) {

	campoEmail, err := Driver.FindElement(selenium.ID, "login_email")
	if err != nil {
		return
	}
	campoEmail.SendKeys(email)

	campoSenha, err := Driver.FindElement(selenium.ID, "login_password")
	if err != nil {
		return
	}
	campoSenha.SendKeys(senha)

	botaoLogin, err := Driver.FindElement(selenium.ByID, "btLogin")

	if err != nil {
		return
	}

	botaoLogin.Click()

	return nil
}

func souAutenticadoComSucesso() (err error) {

	emailMenu, err := Driver.FindElement(selenium, ByClassName, "profile-address")

	if err != nil {
		return
	}

	saida, _ := emailMenu.Text()

	if saida != "eu@papito.io" {
		return fmt.Errorf("Erro ao validar usuario autenticado")
	}

	return nil
}

func devoVerASeguinteMensagem(mensagem string) (err error) {

	divAlerta, err := Driver.FindElement(selenium.ByClassName, "alert-login")

	if err != nil {
		return
	}

	saida, _ := divAlerta.Text()

	if saida != mensagem {
		return fmt.Errorf("Esperava: %v - Obtido: %v", mensagem, saida)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei a pagina principal$`, queEuAcesseiAPaginaPrincipal)
	s.Step(`^faco login com "([^"]*)" e "([^"]*)"$`, faoLoginComE)
	s.Step(`^sou autenticado com sucesso$`, souAutenticadoComSucesso)
	s.Step(`^devo ver a seguinte mensagem "([^"]*)"$`, devoVerASeguinteMensagem)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error) {
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		shot, _ := Driver.Screenshot()

		support.SaveImage(shot, fileName)

		Driver.Quit()

	})

}
