package location

var (
	JKT UNLcode = "JKT"
	JWA UNLcode = "JWA"
	LOM UNLcode = "LOM"
	MAK UNLcode = "MAK"
	MES UNLcode = "MES"
	JAB UNLcode = "JAB"
	JAT UNLcode = "JAT"
	BAL UNLcode = "BAL"
	BAN UNLcode = "BAN"
)

var (
	Jakarta    = &Location{JKT, "Jakarta"}
	JawaTimur  = &Location{JWA, "JawaTimur"}
	Lombok     = &Location{LOM, "Lombok"}
	Makassar   = &Location{MAK, "Makassar"}
	Medan      = &Location{MES, "Medan"}
	JawaBarat  = &Location{JAB, "JawaBarat"}
	JawaTengah = &Location{JAT, "JawaTengah"}
	Bali       = &Location{BAL, "Bali"}
	Bandung    = &Location{BAN, "Bandung"}
)
