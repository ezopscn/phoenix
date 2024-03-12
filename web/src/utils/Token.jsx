// 获取 Token
export const GetToken = () => {
  // 获取 Token 过期时间
  let tokenExpire = sessionStorage.getItem("token-expire");
  if (tokenExpire) {
    // 跟当前时间对比，判断是否过期
    let now = new Date().getTime();
    let timestamp = Date.parse(tokenExpire);
    if (now < timestamp) {
      return sessionStorage.getItem("token");
    }
  }
  // Token 过期或者没有 Token，直接移除这两个 Key，并返回空
  sessionStorage.removeItem("token-expire");
  sessionStorage.removeItem("token");
  return null;
};

// 设置 Token 和过期时间
export const SetToken = (token, expire) => {
  sessionStorage.setItem("token", token);
  sessionStorage.setItem("token-expire", expire);
};
