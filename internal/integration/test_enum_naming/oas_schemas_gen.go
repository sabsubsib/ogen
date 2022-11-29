// Code generated by ogen, DO NOT EDIT.

package api

// Ref: #/components/schemas/PascalExceptionStrat
type PascalExceptionStrat string

const (
	PascalExceptionStrat1      PascalExceptionStrat = "1"
	PascalExceptionStratMinus2 PascalExceptionStrat = "-2"
)

// Ref: #/components/schemas/PascalSpecialStrat
type PascalSpecialStrat string

const (
	PascalSpecialStrat2Plus2  PascalSpecialStrat = "2+2"
	PascalSpecialStrat2Minus2 PascalSpecialStrat = "2-2"
)

// Ref: #/components/schemas/PascalStrat
type PascalStrat string

const (
	PascalStratInSync    PascalStrat = "in-sync"
	PascalStratOutOfSync PascalStrat = "out-of-sync"
)

type ProbeLivenessOK struct {
	VeryBadEnum          VeryBadEnum          `json:"VeryBadEnum"`
	PascalStrat          PascalStrat          `json:"PascalStrat"`
	PascalSpecialStrat   PascalSpecialStrat   `json:"PascalSpecialStrat"`
	PascalExceptionStrat PascalExceptionStrat `json:"PascalExceptionStrat"`
}

// GetVeryBadEnum returns the value of VeryBadEnum.
func (s *ProbeLivenessOK) GetVeryBadEnum() VeryBadEnum {
	return s.VeryBadEnum
}

// GetPascalStrat returns the value of PascalStrat.
func (s *ProbeLivenessOK) GetPascalStrat() PascalStrat {
	return s.PascalStrat
}

// GetPascalSpecialStrat returns the value of PascalSpecialStrat.
func (s *ProbeLivenessOK) GetPascalSpecialStrat() PascalSpecialStrat {
	return s.PascalSpecialStrat
}

// GetPascalExceptionStrat returns the value of PascalExceptionStrat.
func (s *ProbeLivenessOK) GetPascalExceptionStrat() PascalExceptionStrat {
	return s.PascalExceptionStrat
}

// SetVeryBadEnum sets the value of VeryBadEnum.
func (s *ProbeLivenessOK) SetVeryBadEnum(val VeryBadEnum) {
	s.VeryBadEnum = val
}

// SetPascalStrat sets the value of PascalStrat.
func (s *ProbeLivenessOK) SetPascalStrat(val PascalStrat) {
	s.PascalStrat = val
}

// SetPascalSpecialStrat sets the value of PascalSpecialStrat.
func (s *ProbeLivenessOK) SetPascalSpecialStrat(val PascalSpecialStrat) {
	s.PascalSpecialStrat = val
}

// SetPascalExceptionStrat sets the value of PascalExceptionStrat.
func (s *ProbeLivenessOK) SetPascalExceptionStrat(val PascalExceptionStrat) {
	s.PascalExceptionStrat = val
}

// Ref: #/components/schemas/VeryBadEnum
type VeryBadEnum string

const (
	VeryBadEnum_0   VeryBadEnum = " "
	VeryBadEnum_1   VeryBadEnum = "!"
	VeryBadEnum_2   VeryBadEnum = "\""
	VeryBadEnum_3   VeryBadEnum = "#"
	VeryBadEnum_4   VeryBadEnum = "$"
	VeryBadEnum_5   VeryBadEnum = "%"
	VeryBadEnum_6   VeryBadEnum = "&"
	VeryBadEnum_7   VeryBadEnum = "'"
	VeryBadEnum_8   VeryBadEnum = "("
	VeryBadEnum_9   VeryBadEnum = ")"
	VeryBadEnum_10  VeryBadEnum = "*"
	VeryBadEnum_11  VeryBadEnum = "+"
	VeryBadEnum_12  VeryBadEnum = ","
	VeryBadEnum_13  VeryBadEnum = "-"
	VeryBadEnum_14  VeryBadEnum = "."
	VeryBadEnum_15  VeryBadEnum = "/"
	VeryBadEnum_16  VeryBadEnum = "0"
	VeryBadEnum_17  VeryBadEnum = "1"
	VeryBadEnum_18  VeryBadEnum = "2"
	VeryBadEnum_19  VeryBadEnum = "3"
	VeryBadEnum_20  VeryBadEnum = "4"
	VeryBadEnum_21  VeryBadEnum = "5"
	VeryBadEnum_22  VeryBadEnum = "6"
	VeryBadEnum_23  VeryBadEnum = "7"
	VeryBadEnum_24  VeryBadEnum = "8"
	VeryBadEnum_25  VeryBadEnum = "9"
	VeryBadEnum_26  VeryBadEnum = ":"
	VeryBadEnum_27  VeryBadEnum = ";"
	VeryBadEnum_28  VeryBadEnum = "<"
	VeryBadEnum_29  VeryBadEnum = "="
	VeryBadEnum_30  VeryBadEnum = ">"
	VeryBadEnum_31  VeryBadEnum = "?"
	VeryBadEnum_32  VeryBadEnum = "@"
	VeryBadEnum_33  VeryBadEnum = "A"
	VeryBadEnum_34  VeryBadEnum = "B"
	VeryBadEnum_35  VeryBadEnum = "C"
	VeryBadEnum_36  VeryBadEnum = "D"
	VeryBadEnum_37  VeryBadEnum = "E"
	VeryBadEnum_38  VeryBadEnum = "F"
	VeryBadEnum_39  VeryBadEnum = "G"
	VeryBadEnum_40  VeryBadEnum = "H"
	VeryBadEnum_41  VeryBadEnum = "I"
	VeryBadEnum_42  VeryBadEnum = "J"
	VeryBadEnum_43  VeryBadEnum = "K"
	VeryBadEnum_44  VeryBadEnum = "L"
	VeryBadEnum_45  VeryBadEnum = "M"
	VeryBadEnum_46  VeryBadEnum = "N"
	VeryBadEnum_47  VeryBadEnum = "O"
	VeryBadEnum_48  VeryBadEnum = "P"
	VeryBadEnum_49  VeryBadEnum = "Q"
	VeryBadEnum_50  VeryBadEnum = "R"
	VeryBadEnum_51  VeryBadEnum = "S"
	VeryBadEnum_52  VeryBadEnum = "T"
	VeryBadEnum_53  VeryBadEnum = "U"
	VeryBadEnum_54  VeryBadEnum = "V"
	VeryBadEnum_55  VeryBadEnum = "W"
	VeryBadEnum_56  VeryBadEnum = "X"
	VeryBadEnum_57  VeryBadEnum = "Y"
	VeryBadEnum_58  VeryBadEnum = "Z"
	VeryBadEnum_59  VeryBadEnum = "["
	VeryBadEnum_60  VeryBadEnum = "\\"
	VeryBadEnum_61  VeryBadEnum = "]"
	VeryBadEnum_62  VeryBadEnum = "^"
	VeryBadEnum_63  VeryBadEnum = "_"
	VeryBadEnum_64  VeryBadEnum = "`"
	VeryBadEnum_65  VeryBadEnum = "a"
	VeryBadEnum_66  VeryBadEnum = "b"
	VeryBadEnum_67  VeryBadEnum = "c"
	VeryBadEnum_68  VeryBadEnum = "d"
	VeryBadEnum_69  VeryBadEnum = "e"
	VeryBadEnum_70  VeryBadEnum = "f"
	VeryBadEnum_71  VeryBadEnum = "g"
	VeryBadEnum_72  VeryBadEnum = "h"
	VeryBadEnum_73  VeryBadEnum = "i"
	VeryBadEnum_74  VeryBadEnum = "j"
	VeryBadEnum_75  VeryBadEnum = "k"
	VeryBadEnum_76  VeryBadEnum = "l"
	VeryBadEnum_77  VeryBadEnum = "m"
	VeryBadEnum_78  VeryBadEnum = "n"
	VeryBadEnum_79  VeryBadEnum = "o"
	VeryBadEnum_80  VeryBadEnum = "p"
	VeryBadEnum_81  VeryBadEnum = "q"
	VeryBadEnum_82  VeryBadEnum = "r"
	VeryBadEnum_83  VeryBadEnum = "s"
	VeryBadEnum_84  VeryBadEnum = "t"
	VeryBadEnum_85  VeryBadEnum = "u"
	VeryBadEnum_86  VeryBadEnum = "v"
	VeryBadEnum_87  VeryBadEnum = "w"
	VeryBadEnum_88  VeryBadEnum = "x"
	VeryBadEnum_89  VeryBadEnum = "y"
	VeryBadEnum_90  VeryBadEnum = "z"
	VeryBadEnum_91  VeryBadEnum = "{"
	VeryBadEnum_92  VeryBadEnum = "|"
	VeryBadEnum_93  VeryBadEnum = "}"
	VeryBadEnum_94  VeryBadEnum = "~"
	VeryBadEnum_95  VeryBadEnum = "id"
	VeryBadEnum_96  VeryBadEnum = "id desc"
	VeryBadEnum_97  VeryBadEnum = "classifyAs"
	VeryBadEnum_98  VeryBadEnum = "classifyAs desc"
	VeryBadEnum_99  VeryBadEnum = "senderEmailAddress"
	VeryBadEnum_100 VeryBadEnum = "senderEmailAddress desc"
	VeryBadEnum_101 VeryBadEnum = "childFolders"
	VeryBadEnum_102 VeryBadEnum = "messageRules"
	VeryBadEnum_103 VeryBadEnum = "messages"
	VeryBadEnum_104 VeryBadEnum = "multiValueExtendedProperties"
	VeryBadEnum_105 VeryBadEnum = "singleValueExtendedProperties"
)