//Package logf comment
// This file war generated by tars2go 1.1
// Generated from LogF.tars
package logf

import (
	"fmt"
	"gitee.com/bee-circle/tarsgo/tars/protocol/codec"
)

//LogInfo strcut implement
type LogInfo struct {
	Appname           string `json:"appname"`
	Servername        string `json:"servername"`
	SFilename         string `json:"sFilename"`
	SFormat           string `json:"sFormat"`
	Setdivision       string `json:"setdivision"`
	BHasSufix         bool   `json:"bHasSufix"`
	BHasAppNamePrefix bool   `json:"bHasAppNamePrefix"`
	BHasSquareBracket bool   `json:"bHasSquareBracket"`
	SConcatStr        string `json:"sConcatStr"`
	SSepar            string `json:"sSepar"`
	SLogType          string `json:"sLogType"`
}

func (st *LogInfo) resetDefault() {
	st.BHasSufix = true
	st.BHasAppNamePrefix = true
	st.BHasSquareBracket = false
	st.SConcatStr = "_"
	st.SSepar = "|"
	st.SLogType = ""
}

//ReadFrom reads  from _is and put into struct.
func (st *LogInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Appname, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Servername, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SFilename, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SFormat, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Setdivision, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BHasSufix, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BHasAppNamePrefix, 6, false)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BHasSquareBracket, 7, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SConcatStr, 8, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SSepar, 9, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SLogType, 10, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *LogInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require LogInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *LogInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Appname, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Servername, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SFilename, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SFormat, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Setdivision, 4)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BHasSufix, 5)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BHasAppNamePrefix, 6)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BHasSquareBracket, 7)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SConcatStr, 8)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SSepar, 9)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SLogType, 10)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *LogInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
