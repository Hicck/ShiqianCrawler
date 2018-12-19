package utils

import "regexp"

var TimerReg = regexp.MustCompile(`[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9] [0-9][0-9]:[0-9][0-9]:[0-9][0-9]`)
var TitleRg1 = regexp.MustCompile(`[[:^ascii:]]`)
var TitleRg2 = regexp.MustCompile(`[[:^ascii:]|[0-9]+`)
var EditReg = regexp.MustCompile(`[[:^ascii:]|[0-9]+`)
var IdReg = regexp.MustCompile(`uploadfile/([0-9]+)/([0-9]+)/`)
var CheckId = regexp.MustCompile(`[0-9]+`)
