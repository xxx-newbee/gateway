// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package chat

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/xxx-newbee/chat/chat"
	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatRequest, w http.ResponseWriter) (*types.BaseResponse, error) {
	if req.Prompt == "" || req.Model == "" {
		return &types.BaseResponse{
			Code: 200,
			Msg:  "prompt or model empty",
			Data: nil,
		}, nil
	}

	if req.Stream == true {
		streamResp, err := l.svcCtx.ChatRpc.StreamChat(l.ctx, &chat.ChatReq{
			Prompt: req.Prompt,
			Model:  req.Model,
		})
		if err != nil {
			l.Logger.Errorf("[ChatLogic] StreamChat error: %v", err)
		}
		for {
			resp, err := streamResp.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Fprintf(w, "data: {\"type\":\"done\",\"message\":\"Stream completed\"}\n")
					if flusher, ok := w.(http.Flusher); ok {
						flusher.Flush()
					}
					break
				}
				// 发送错误信息
				fmt.Fprintf(w, "data: {\"type\":\"error\",\"message\":\"%s\"}\n", err.Error())
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
				break
			}
			// 构造 SSE 数据
			if resp.Msg != "" {
				fmt.Fprintf(w, "data: {\"type\":\"error\",\"message\":\"%s\"}\n", resp.Msg)
			} else if resp.Response != "" {
				fmt.Fprintf(w, "data: {\"type\":\"chunk\",\"content\":\"%s\"}\n", resp.Response)
			}

			if resp.Done {
				fmt.Fprintf(w, "data: {\"type\":\"done\",\"message\":\"Final completion\"}\n")
			}

			// 确保数据立即发送
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}

			if resp.Done {
				break
			}

		}
		return nil, nil
	} else {
		resp, err := l.svcCtx.ChatRpc.Chat(l.ctx, &chat.ChatReq{
			Prompt: req.Prompt,
			Model:  req.Model,
		})
		if err != nil {
			return &types.BaseResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}

		return &types.BaseResponse{
			Code: 200,
			Data: resp,
		}, nil

	}

}
