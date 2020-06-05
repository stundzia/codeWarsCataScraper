package models

type KataPool struct {
	Katas []*Kata
}

type Kata struct {
	Title string
	Url string
	Kyu int
	Langs []int
}

func (katas *KataPool) FindKatasByKyu(kyu int) (res []*Kata) {
	for _, k := range katas.Katas {
		if k.Kyu == kyu {
			res = append(res, k)
		}
	}
	return res
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func (kata *Kata) SupportsLangs(langs []int) bool {
	for _, l := range langs {
		if !contains(kata.Langs, l) {
			return false
		}
	}
	return true
}

func (katas *KataPool) FindKatasByLangList(langList []int) (res []*Kata) {
	for _, k := range katas.Katas {
		if !k.SupportsLangs(langList) {
			continue
		}
		res = append(res, k)
	}
	return res
}

func (katas *KataPool) GetKataByTitle(title string) *Kata {
	for _, kata := range katas.Katas {
		if kata.Title == title {
			return kata
		}
	}
	return nil
}

func (katas *KataPool) GetKataByKyuRangeAndLanguages(kyuMin int, kyuMax int, langs []int) (res []*Kata)  {
	for _, kata := range katas.Katas {
		if kata.Kyu < kyuMin || kata.Kyu > kyuMax || !kata.SupportsLangs(langs) {
			continue
		}
		res = append(res, kata)
	}
	return res
}