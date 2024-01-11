package tags

import textrank "github.com/DavidBelicza/TextRank"

func GetTopPhrasesFromText(text string) []string {
	// generate list of key phrases using TextRank algorithm
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()

	tr.Populate(text, language, rule)
	tr.Ranking(algorithmDef)

	// Get top n phrases with weight over 0.5
	rankedPhrases := textrank.FindPhrases(tr)
	res := []string{}
	for i := 0; i <= 5; i++ {
		rp := rankedPhrases[i]
		if rp.Weight >= 0.5 {
			phrase := rp.Left + " " + rp.Right
			res = append(res, phrase)
		}
	}
	return res
}
