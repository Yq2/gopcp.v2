package module

// Type 代表组件的类型。
type Type string

// 当前认可的组件类型的常量。
const (
	// TYPE_DOWNLOADER 代表下载器。
	TypeDownloader Type = "downloader"
	// TYPE_ANALYZER 代表分析器。
	TypeAnalyzer Type = "analyzer"
	// TYPE_PIPELINE 代表条目处理管道。
	TypePipeline Type = "pipeline"
)

// legalTypeLetterMap 代表合法的组件类型-字母的映射。
var legalTypeLetterMap = map[Type]string{
	TypeDownloader: "D",
	TypeAnalyzer:   "A",
	TypePipeline:   "P",
}

// legalLetterTypeMap 代表合法的字母-组件类型的映射。
var legalLetterTypeMap = map[string]Type{
	"D": TypeDownloader,
	"A": TypeAnalyzer,
	"P": TypePipeline,
}

// CheckType 用于判断组件实例的类型是否匹配。
func CheckType(moduleType Type, module Module) bool {
	if moduleType == "" || module == nil {
		return false
	}
	switch moduleType {
	case TypeDownloader:
		if _, ok := module.(Downloader); ok {
			return true
		}
	case TypeAnalyzer:
		if _, ok := module.(Analyzer); ok {
			return true
		}
	case TypePipeline:
		if _, ok := module.(Pipeline); ok {
			return true
		}
	}
	return false
}

// LegalType 用于判断给定的组件类型是否合法。
func LegalType(moduleType Type) bool {
	if _, ok := legalTypeLetterMap[moduleType]; ok {
		return true
	}
	return false
}

// GetType 用于获取组件的类型。
// 若给定的组件ID不合法则第一个结果值会是false。
func GetType(mid MID) (bool, Type) {
	parts, err := SplitMID(mid)
	if err != nil {
		return false, ""
	}
	mt, ok := legalLetterTypeMap[parts[0]]
	return ok, mt
}

// getLetter 用于获取组件类型的字母代号。
func getLetter(moduleType Type) (bool, string) {
	var letter string
	var found bool
	for l, t := range legalLetterTypeMap {
		if t == moduleType {
			letter = l
			found = true
			break
		}
	}
	return found, letter
}

// typeToLetter 用于根据给定的组件类型获得其字母代号。
// 若给定的组件类型不合法，则第一个结果值会是false。
func typeToLetter(moduleType Type) (bool, string) {
	switch moduleType {
	case TypeDownloader:
		return true, "D"
	case TypeAnalyzer:
		return true, "A"
	case TypePipeline:
		return true, "P"
	default:
		return false, ""
	}
}

// letterToType 用于根据字母代号获得对应的组件类型。
// 若给定的字母代号不合法，则第一个结果值会是false。
func letterToType(letter string) (bool, Type) {
	switch letter {
	case "D":
		return true, TypeDownloader
	case "A":
		return true, TypeAnalyzer
	case "P":
		return true, TypePipeline
	default:
		return false, ""
	}
}
