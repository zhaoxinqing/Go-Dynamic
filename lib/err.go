package lib

import (
	"net/http"
)

const (
	UserRegistered                    = 100101 // 用户已注册
	UserPassword                      = 100102 // 用户登录密码错误
	UserNotRegistered                 = 100103 // 您还未注册，请先注册
	UserRegisteredError               = 100104 // 用户注册失败
	UserPasswordError                 = 100105 // 用户重制密码失败
	UserEmailFrequentError            = 100106 // 邮件发送频繁
	UserEmailSendError                = 100107 // 邮件发送频繁
	EmailCodeFailure                  = 100108 // 验证码已失效
	EmailCodeError                    = 100109 // 验证码不正确
	EmailCodeNotMatch                 = 100110 // 验证码和邮件不匹配
	CollectionNotExist                = 100111 // 该藏品不存在
	CollectionSoldOut                 = 100112 // 该藏品已售完
	CollectionNotEnough               = 100113 // 该藏品数量不足
	OrderBuyNumberNull                = 100114 // 未输入购买数量
	NoAppointment                     = 100115 // 未进行预约
	ExistUnpaidOrder                  = 100116 // 有未支付的订单
	CreateOrderFailed                 = 100117 // 创建订单失败
	PayOrderFailed                    = 100118 // 一级市场支付订单失败
	OrderNotExist                     = 100119 // 订单不存在
	OrderIsFromMarket                 = 100200 // 订单来自二级市场
	OrderStatusError                  = 100201 // 订单状态不正确
	TradePasswordEmpty                = 100202 // 交易密码不能为空
	NotSetTradePassword               = 100203 // 未设置交易密码
	UserTradePasswordError            = 100204 // 用户交易密码错误
	TooManyTradePassword              = 100205 // 今日输入密码已达上限，不可输入
	UserNoMoney                       = 100206 // 余额不足
	UserPaymentFailed                 = 100207 // 支付失败
	CollectionNotUser                 = 100208 // 藏品不属于某个用户
	CollectionConsignmentIng          = 100209 // 当前藏品正在寄售，不可重复寄售
	CollectionConsignmentError        = 100210 // 当前藏品寄售失败
	CollectionConsignmentNumberError  = 100211 // 当前藏品数量不足
	ConsignmentOrderNotFound          = 100212 // 寄售订单不存在
	ConsignmentOrderStatusError       = 100213 // 寄售订单状态不正确
	ConsignmentCancelError            = 100214 // 藏品取消寄售失败
	ConsignmentCancelSuccess          = 100215 // 藏品取消寄售成功
	ConsignmentOrderNotUser           = 100216 // 当前寄售订单不属于你
	OrderPaid                         = 100217 // 订单已交易完成
	OrderLocking                      = 100219 // 当前正在交易中，请稍后重试
	NotBuyYourself                    = 100220 // 暂不支持购买自己的寄售订单
	ConsignmentNoUp                   = 100221 // 当前藏品无人挂单
	OrderNotCancel                    = 100222 // 订单已支付，不允许取消
	OrderClosed                       = 100223 // 订单已关闭
	OrderCancel                       = 100224 // 取消购买失败
	ConsignmentNotExist               = 100226 // 寄售不存在
	OrderFailedNotExist               = 100227 // 订单不存在
	TwoPasswordsDifferent             = 100228 // 两次输入的密码不一致
	UserNotVerified                   = 100229 // 用户未进行实名认证
	WithdrawLowerLimited              = 100230 // 低于最小提现金额
	WithdrawUpperLimited              = 100231 // 超过提现限制（一天一次）
	CollectionExpired                 = 100233 // 发售期已过
	CollectionUnStart                 = 100234 // 暂未开售
	CollectionUnReserve               = 100235 // 未预约
	CollectionCurrentTimeWhitelistBuy = 100236 // 当前时间仅限白名单用户购买
	CollectionUpperLimit              = 100237 // 购买已达到上限 白名单限购已达上限，正式开售可再次购买
	CollectionReserveExpired          = 100238 // 已过预约时间
	CollectionReserveStop             = 100239 // 预约人数已达上限，暂停预约
	CollectionReserveFailed           = 100240 // 预约失败
	CollectionReserveExists           = 100241 // 重复预约
	UnLogin                           = 100243 // 请先登录
	SellPriceLow                      = 100244 // 售出价格不得低于
	WalletAddressNotFound             = 100245 // 钱包地址不存在
	AmountExceedsLimit                = 100246 // 提现金额超出限制
	UploadLimit                       = 100247 // 上传次数达到上限
	ConsignmentUnStart                = 100248 // 未到寄售时间
	ChangePasswordNotTrad             = 100249 // 24h后才能发起新的交易
	UserVerificationExist             = 100250 // 该证件已经被使用
	UserVerificationLimit             = 100251 // 审核失败次数过多，无法继续认证
	UserWalletBalanceMore             = 100252 // 账户资产大于等于 10
	UserCancelAccountFail             = 100253 // 注销失败
	InstitutionNotExist               = 100254 // 机构不存在
	WithdrawAmountNotEnough           = 100256 // 超出可提数量
	OnListingNotWithdraw              = 100257 // 寄售中，无法提取
	CollectibleWithdrawFailed         = 100258 // 提现失败，请稍后重试
	IncorrectDBRowsAffected           = 100259 // 错误的db执行影响记录数
	ChangePasswordNotTrade            = 100260 // 重置密码24h后才能发起新的交易

	Unauthorized        = 401001 // Unauthorized   未授权
	CodeInvalidParams   = 400001 // 4xx 客户端错误
	CodeNoRoute         = 400102
	CollectibleNotFound = 400103 // 藏品系列未找到

	CodeOutOfLimit = 429000 // 限流

	CodeInternalServer = 500100 // 服务器内部错误
	CodeFailedService  = 500101 // 5xx 服务端错误
	CodeFailedDb       = 500103
	InternalSystemBusy = 500106

	CodeGoogleRecaptchaFailed = 600102 // google 人机验证失败
)

const (
	// MsgUnauthorized 400 类型的系统统一错误提示
	MsgUnauthorized = "Unauthorized"
	// MsgFailedDb 500 服务统一错误提示
	MsgFailedDb = "数据异常"
)

type CodeError struct {
	httpCode int
	errCode  int
	errMsg   string
}

func (e *CodeError) GetHttpCode() int {
	if e.httpCode == 0 {
		return http.StatusOK
	}
	return e.httpCode
}

func (e *CodeError) GetErrCode() int {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func New(errCode int, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

// CustomErrorMessage 接口参数校验
func CustomErrorMessage(msg string) *CodeError {
	return &CodeError{errCode: CodeInvalidParams, errMsg: msg}
}

// ExpiredToken token 失效  Unauthorized 401 接口返回
func ExpiredToken() *CodeError {
	return New(Unauthorized, MsgUnauthorized)
}

// FailedDB 数据库异常
func FailedDB(err error) *CodeError {
	return New(CodeFailedDb, MsgFailedDb+":"+err.Error())
}
