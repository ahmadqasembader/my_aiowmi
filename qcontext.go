package main

type QContext struct {
	q string
	conn Connection
	proto Protocol
	flags int
	timeout int
	skip_optimize bool
}

func (qc *QContext) NewQContext(q string, conn Connection, proto Protocol, flags, timeout int, skip_optimize bool) *QContext{
	qc.q = q
	qc.conn = conn
	qc.proto = proto
	qc.flags = flags
	qc.timeout = timeout
	qc.skip_optimize = skip_optimize
	
	// TODO
	// qc.class_parts
	// qc.interface
	// qc.next

	return qc
}

