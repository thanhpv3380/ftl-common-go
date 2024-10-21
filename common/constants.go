package common

const COMMON_JS_VERSION = "1.1"

var PAGINATION = struct {
	PAGE_SIZE       int
	PAGE_NUM        int
	PAGE_SIZE_LIMIT int
}{
	PAGE_SIZE:       10,
	PAGE_NUM:        1,
	PAGE_SIZE_LIMIT: 1000,
}

type LANGUAGE string

const (
	VI LANGUAGE = "vi"
	EN LANGUAGE = "en"
)

type PLATFORM string

const (
	MTS PLATFORM = "MTS"
	WTS PLATFORM = "WTS"
)

type SUB_TYPE string

const (
	EQUITY      SUB_TYPE = "EQUITY"
	DERIVATIVES SUB_TYPE = "DERIVATIVES"
)

type INVEST_TYPE string

const (
	NORMAL INVEST_TYPE = "NORMAL"
	MARGIN INVEST_TYPE = "MARGIN"
)
