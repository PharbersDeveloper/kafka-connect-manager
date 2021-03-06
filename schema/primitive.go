// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     ConnectRequest.avsc
 *     ConnectResponse.avsc
 *     MonitorRequest.avsc
 *     MonitorResponse.avsc
 *     TmAggRequest.avsc
 *     TmAggResponse.avsc
 */

package schema

import (
	"io"
)

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func writeConnectRequest(r *ConnectRequest, w io.Writer) error {
	var err error
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeString(r.Tag, w)
	if err != nil {
		return err
	}
	err = writeString(r.SourceConfig, w)
	if err != nil {
		return err
	}
	err = writeString(r.SinkConfig, w)
	if err != nil {
		return err
	}

	return nil
}
func writeConnectResponse(r *ConnectResponse, w io.Writer) error {
	var err error
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeLong(r.Progress, w)
	if err != nil {
		return err
	}
	err = writeString(r.Status, w)
	if err != nil {
		return err
	}
	err = writeString(r.Message, w)
	if err != nil {
		return err
	}

	return nil
}

func writeInt(r int32, w io.Writer) error {
	downShift := uint32(31)
	encoded := uint64((uint32(r) << 1) ^ uint32(r>>downShift))
	const maxByteSize = 5
	return encodeInt(w, maxByteSize, encoded)
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeMonitorRequest2(r *MonitorRequest2, w io.Writer) error {
	var err error
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeString(r.ConnectorName, w)
	if err != nil {
		return err
	}
	err = writeString(r.SourceTopic, w)
	if err != nil {
		return err
	}
	err = writeString(r.RecallTopic, w)
	if err != nil {
		return err
	}
	err = writeString(r.Strategy, w)
	if err != nil {
		return err
	}

	return nil
}
func writeMonitorResponse2(r *MonitorResponse2, w io.Writer) error {
	var err error
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeString(r.ConnectorName, w)
	if err != nil {
		return err
	}
	err = writeLong(r.Progress, w)
	if err != nil {
		return err
	}
	err = writeString(r.Error, w)
	if err != nil {
		return err
	}

	return nil
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

func writeTmAggRequest(r *TmAggRequest, w io.Writer) error {
	var err error
	err = writeString(r.RequestId, w)
	if err != nil {
		return err
	}
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeString(r.ProposalId, w)
	if err != nil {
		return err
	}
	err = writeString(r.ProjectId, w)
	if err != nil {
		return err
	}
	err = writeString(r.PeriodId, w)
	if err != nil {
		return err
	}
	err = writeInt(r.Phase, w)
	if err != nil {
		return err
	}
	err = writeString(r.Strategy, w)
	if err != nil {
		return err
	}

	return nil
}
func writeTmAggResponse(r *TmAggResponse, w io.Writer) error {
	var err error
	err = writeString(r.RequestId, w)
	if err != nil {
		return err
	}
	err = writeString(r.JobId, w)
	if err != nil {
		return err
	}
	err = writeString(r.Status, w)
	if err != nil {
		return err
	}
	err = writeString(r.Message, w)
	if err != nil {
		return err
	}

	return nil
}
