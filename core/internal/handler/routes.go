// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"gcloud/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/mail/code/send/register",
				Handler: MailCodeSendRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/share/basic/detail",
				Handler: ShareBasicDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/popular/share/list",
				Handler: PopularShareListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/share/statistics",
				Handler: ShareStatisticsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/register/count",
				Handler: RegisterCountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/public/file/list",
				Handler: PublicFileListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/list",
				Handler: PostsListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/comment",
				Handler: PostsCommentsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/posts/detail",
				Handler: PostsDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/gongde/update",
				Handler: GongDeUpdateHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/public/file/save",
					Handler: PublicFileSaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/ws/message",
					Handler: WebsocketMessageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload",
					Handler: FileUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: UserUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: UserRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/file/list",
					Handler: UserFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/name/update",
					Handler: UserFileNameUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/folder/create",
					Handler: UserFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file/delete",
					Handler: UserFileDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/public/file/name/update",
					Handler: PublicFileNameUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/public/folder/create",
					Handler: PublicFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/public/file/delete",
					Handler: PublicFileDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/user/file/move",
					Handler: UserFileMoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/download",
					Handler: FileDownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/create",
					Handler: ShareBasicCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share/basic/save",
					Handler: ShareBasicSaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/share/list",
					Handler: UserShareListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/refresh/authorization",
					Handler: RefreshAuthorizationHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/detail",
					Handler: UserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/prepare",
					Handler: FileUploadPrepareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/chunk",
					Handler: FileUploadChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/file/upload/chunk/complete",
					Handler: FileUploadChunkCompleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/posts/create",
					Handler: PostsCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/posts/update",
					Handler: PostsUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/posts/delete",
					Handler: PostsDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/posts/comment/create",
					Handler: PostsCommentCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/posts/comment/delete",
					Handler: PostsCommentDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/posts/feedback/create",
					Handler: PostsFeedbackCreateHandler(serverCtx),
				},
			}...,
		),
	)
}
