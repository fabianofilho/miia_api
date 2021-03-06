package models

type Semus struct {
	Keys   []string `json:"keys"`
	Values []string `json:"values"`
}

type SemusOutPayload struct {
	Json map[string]string
}

type SemusSQL struct {
	ID        string `json:"idsemus"`
	IDUser    string `json:"iduser"`
	IDCompany string `json:"idcompany"`
	Result    string `json:"result"`
	Input     string `json:"input"`
}

type SemusRecPayload struct {
	Apgar5      string `json:"apgar5" schema:"apgar5,required"`
	Peso        string `json:"peso" schema:"peso,required"`
	IdadeMae    string `json:"idade_mae" schema:"idade_mae,required"`
	Consultas   string `json:"consultas" schema:"consultas,required"`
	Gestacao    string `json:"gestacao" schema:"gestacao,required"`
	SitConjuMae string `json:"sit_conjugal_mae" schema:"sit_conjugal_mae,required"`
	Sexo        string `json:"sexo" schema:"sexo,required"`
	Anomalia    string `json:"anomalia" schema:"anomalia,required"`
	Gravidez    string `json:"gravidez" schema:"gravidez,required"`
	Parto       string `json:"parto" schema:"parto,required"`
}

func (cc SemusRecPayload) Parse() (s Semus) {
	s.Keys = semusKeys

	s.Values = []string{cc.IdadeMae}
	//s.Values = append(s.Values, anomalia[cc.Anomalia]...)
	s.Values = append(s.Values, gestacao[cc.Gestacao]...)
	s.Values = append(s.Values, consultas[cc.Consultas]...)
	s.Values = append(s.Values, estadocivil[cc.SitConjuMae]...)
	s.Values = append(s.Values, sexo[cc.Sexo]...)
	s.Values = append(s.Values, gravidez[cc.Gravidez]...)
	s.Values = append(s.Values, parto[cc.Parto])
	s.Values = append(s.Values, []string{cc.Apgar5, cc.Peso}...)
	return s
}

func (cc Semus) Parse() (s SemusOutPayload) {
	s.Json = make(map[string]string)
	if len(cc.Keys) == len(cc.Values) {
		for i, key := range cc.Keys {
			s.Json[key] = cc.Values[i]
		}
	}
	return s
}

func (cc SemusSQL) Parse() (s Semus) {
	return s
}

// =====  Constants =====
var semusKeys []string = []string{
	"IDADEMAE_x",
	"GESTACAO_x_1.0", "GESTACAO_x_2.0", "GESTACAO_x_3.0", "GESTACAO_x_4.0", "GESTACAO_x_5.0", "GESTACAO_x_6.0", "GESTACAO_x_9.0",
	"CONSULTAS_1.0", "CONSULTAS_2.0", "CONSULTAS_3.0", "CONSULTAS_4.0", "CONSULTAS_9.0",
	"ESTCIVMAE_1.0", "ESTCIVMAE_2.0", "ESTCIVMAE_3.0", "ESTCIVMAE_4.0", "ESTCIVMAE_5.0",
	"SEXO_x_F", "SEXO_x_I", "SEXO_x_M",
	"GRAVIDEZ_x_1.0", "GRAVIDEZ_x_2.0", "GRAVIDEZ_x_3.0",
	"PARTO_x_1.0",
	"APGAR5", "PESO_x",
}

var sexo map[string][]string = map[string][]string{
	"m": []string{"0", "0", "1"},
	"i": []string{"0", "1", "0"},
	"f": []string{"1", "0", "0"},
}

var anomalia map[string][]string = map[string][]string{
	"1": []string{"1", "0", "0"},
	"2": []string{"0", "1", "0"},
	"9": []string{"0", "0", "1"},
}

var gestacao map[string][]string = map[string][]string{
	"1": []string{"1", "0", "0", "0", "0", "0", "0"},
	"2": []string{"0", "1", "0", "0", "0", "0", "0"},
	"3": []string{"0", "0", "1", "0", "0", "0", "0"},
	"4": []string{"0", "0", "0", "1", "0", "0", "0"},
	"5": []string{"0", "0", "0", "0", "1", "0", "0"},
	"6": []string{"0", "0", "0", "0", "0", "1", "0"},
	"9": []string{"0", "0", "0", "0", "0", "0", "1"},
}

var gravidez map[string][]string = map[string][]string{
	"1": []string{"0", "1", "0"},
	"2": []string{"1", "0", "0"},
	"3": []string{"0", "0", "0"},
}

var consultas map[string][]string = map[string][]string{
	"1": []string{"1", "0", "0", "0", "0"},
	"2": []string{"0", "1", "0", "0", "0"},
	"3": []string{"0", "0", "1", "0", "0"},
	"4": []string{"0", "0", "0", "1", "0"},
	"9": []string{"0", "0", "0", "0", "1"},
}

var parto map[string]string = map[string]string{
	"1": "1",
	"2": "0",
}

var estadocivil map[string][]string = map[string][]string{
	"1": []string{"1", "0", "0", "0", "0"},
	"2": []string{"0", "1", "0", "0", "0"},
	"3": []string{"0", "0", "1", "0", "0"},
	"4": []string{"0", "0", "0", "1", "0"},
	"5": []string{"0", "0", "0", "0", "1"},
}
