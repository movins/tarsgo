//Package statf comment
// This file war generated by tars2go 1.1
// Generated from StatF.tars
package statf

import (
	"fmt"
	"gitee.com/bee-circle/tarsgo/tars/protocol/codec"
)

//StatSampleMsg strcut implement
type StatSampleMsg struct {
	Unid          string `json:"unid"`
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	InterfaceName string `json:"interfaceName"`
	MasterIp      string `json:"masterIp"`
	SlaveIp       string `json:"slaveIp"`
	Depth         int32  `json:"depth"`
	Width         int32  `json:"width"`
	ParentWidth   int32  `json:"parentWidth"`
}

func (st *StatSampleMsg) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *StatSampleMsg) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Unid, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.MasterName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveName, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.InterfaceName, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.MasterIp, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SlaveIp, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Depth, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Width, 7, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ParentWidth, 8, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *StatSampleMsg) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatSampleMsg, but not exist. tag %d", tag)
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
func (st *StatSampleMsg) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Unid, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.MasterName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveName, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.InterfaceName, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.MasterIp, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SlaveIp, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Depth, 6)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Width, 7)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ParentWidth, 8)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *StatSampleMsg) WriteBlock(_os *codec.Buffer, tag byte) error {
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
