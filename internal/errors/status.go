package errors

var (
	UnknownError = Error{1000001, 400, "알 수 없는 오류"}

	ShortCipherBlockSize   = Error{1000401, 500, "Cipher Text 블록 사이즈가 너무 짧습니다"}
	ConversionTextNotFound = Error{1000402, 404, "변환하려는 텍스트가 존재하지 않습니다"}

	CacheDataConfigError = Error{1001006, 409, "캐시 데이터 설정간 오류가 발생하였습니다"}
)
