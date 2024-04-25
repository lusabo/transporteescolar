package initializers

import (
	"backend/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Falha ao conectar no banco de dados")
	}

}

func SyncDB() {
	DB.AutoMigrate(&models.Estado{}, &models.Cidade{}, &models.Bairro{}, &models.Escola{}, &models.Motorista{})
}

func Seed() {
	estados := []models.Estado{
		{Uf: "AC", Nome: "Acre"},
		{Uf: "AL", Nome: "Alagoas"},
		{Uf: "AP", Nome: "Amapá"},
		{Uf: "AM", Nome: "Amazonas"},
		{Uf: "BA", Nome: "Bahia"},
		{Uf: "CE", Nome: "Ceará"},
		{Uf: "DF", Nome: "Distrito Federal"},
		{Uf: "ES", Nome: "Espírito Santo"},
		{Uf: "GO", Nome: "Goiás"},
		{Uf: "MA", Nome: "Maranhão"},
		{Uf: "MT", Nome: "Mato Grosso"},
		{Uf: "MS", Nome: "Mato Grosso do Sul"},
		{Uf: "MG", Nome: "Minas Gerais"},
		{Uf: "PA", Nome: "Pará"},
		{Uf: "PB", Nome: "Paraíba"},
		{Uf: "PR", Nome: "Paraná"},
		{Uf: "PE", Nome: "Pernambuco"},
		{Uf: "PI", Nome: "Piauí"},
		{Uf: "RJ", Nome: "Rio de Janeiro"},
		{Uf: "RN", Nome: "Rio Grande do Norte"},
		{Uf: "RS", Nome: "Rio Grande do Sul"},
		{Uf: "RO", Nome: "Rondônia"},
		{Uf: "RR", Nome: "Roraima"},
		{Uf: "SC", Nome: "Santa Catarina"},
		{Uf: "SP", Nome: "São Paulo"},
		{Uf: "SE", Nome: "Sergipe"},
		{Uf: "TO", Nome: "Tocantins"},
	}

	for _, estado := range estados {
		DB.Create(&estado)
		switch estado.Uf {
		case "AC":
			addCidades(estado, []string{"Rio Branco", "Cruzeiro do Sul", "Sena Madureira"})
		case "AL":
			addCidades(estado, []string{"Maceió", "Arapiraca", "Rio Largo"})
		case "AP":
			addCidades(estado, []string{"Macapá", "Santana", "Laranjal do Jari"})
		case "AM":
			addCidades(estado, []string{"Manaus", "Parintins", "Itacoatiara"})
		case "BA":
			addCidades(estado, []string{"Salvador", "Feira de Santana", "Vitória da Conquista"})
		case "CE":
			addCidades(estado, []string{"Fortaleza", "Caucaia", "Juazeiro do Norte"})
		case "DF":
			addCidades(estado, []string{"Brasília", "Ceilândia", "Gama"})
		case "ES":
			addCidades(estado, []string{"Vitória", "Vila Velha", "Serra"})
		case "GO":
			addCidades(estado, []string{"Goiânia", "Aparecida de Goiânia", "Anápolis"})
		case "MA":
			addCidades(estado, []string{"São Luís", "Imperatriz", "São José de Ribamar"})
		case "MT":
			addCidades(estado, []string{"Cuiabá", "Várzea Grande", "Rondonópolis"})
		case "MS":
			addCidades(estado, []string{"Campo Grande", "Dourados", "Três Lagoas"})
		case "MG":
			addCidades(estado, []string{"Belo Horizonte", "Uberlândia", "Contagem"})
		case "PA":
			addCidades(estado, []string{"Belém", "Ananindeua", "Santarém"})
		case "PB":
			addCidades(estado, []string{"João Pessoa", "Campina Grande", "Santa Rita"})
		case "PR":
			addCidades(estado, []string{"Curitiba", "Londrina", "Maringá"})
		case "PE":
			addCidades(estado, []string{"Recife", "Jaboatão dos Guararapes", "Olinda"})
		case "PI":
			addCidades(estado, []string{"Teresina", "Parnaíba", "Picos"})
		case "RJ":
			addCidades(estado, []string{"Rio de Janeiro", "São Gonçalo", "Duque de Caxias"})
		case "RN":
			addCidades(estado, []string{"Natal", "Mossoró", "Parnamirim"})
		case "RS":
			addCidades(estado, []string{"Porto Alegre", "Caxias do Sul", "Canoas"})
		case "RO":
			addCidades(estado, []string{"Porto Velho", "Ji-Paraná", "Ariquemes"})
		case "RR":
			addCidades(estado, []string{"Boa Vista", "Rorainópolis", "Caracaraí"})
		case "SC":
			addCidades(estado, []string{"Florianópolis", "Joinville", "São José"})
		case "SP":
			addCidades(estado, []string{"São Paulo", "Guarulhos", "Campinas"})
		case "SE":
			addCidades(estado, []string{"Aracaju", "Nossa Senhora do Socorro", "Lagarto"})
		case "TO":
			addCidades(estado, []string{"Palmas", "Araguaína", "Gurupi"})
		default:
			// Adicione um tratamento de erro caso não haja dados disponíveis para o estado
			fmt.Printf("Dados de cidades não encontrados para %s\n", estado.Uf)
		}
	}
}

func addCidades(estado models.Estado, nomes []string) {
	for _, nome := range nomes {
		cidade := models.Cidade{Nome: nome, EstadoID: estado.ID}
		DB.Create(&cidade)
	}
}
