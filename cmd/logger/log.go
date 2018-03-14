package logger

import "context"

const contextKey = "reqInfo"

type KeyVal struct {
	Key string
	Val string
}

type ReqInfo struct {
	RemoteHost string
	UserAgent  string
	RequestID  string
	API        string
	BucketName string
	ObjectName string
	Tags       []KeyVal
}

func (r *ReqInfo) AppendTags(key string, val string) {
	if r == nil {
		return
	}
	r.Tags = append(r.Tags, KeyVal{key, val})
}

func ContextSet(ctx context.Context, req *ReqInfo) context.Context {
	return context.WithValue(ctx, contextKey, req)
}

func ContextGet(ctx context.Context) *ReqInfo {
	r, ok := ctx.Value(contextKey).(*ReqInfo)
	if ok {
		return r
	}
	return nil
}
