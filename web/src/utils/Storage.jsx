// 保存数据，支持设置有效期（单位秒）
function SetLocalStorageItem(key, value, expire) {
  // 当前时间
  let now = Date.now();
  // 判断传入的有效期是否为数值或有效，如果不合理，则永不过期
  if (isNaN(expire) || expire < 0) {
    expire = 0;
  } else {
    expire = expire * 1000;
  }
  // 将值封装成对象
  let obj = {
    data: value,
    time: now,
    expire: expire,
  };
  // 注意，localStorage 不能直接存储对象类型，需要先用 JSON.stringify() 将其转换成字符串，取值时再通过 JSON.parse() 转换回来
  localStorage.setItem(key, JSON.stringify(obj));
}

// 获取数据
function GetLocalStorageItem(key) {
  let val = localStorage.getItem(key);

  // 如果没有值就直接返回 null
  if (!val) {
    return null;
  }

  // 判断是否过期
  val = JSON.parse(val);
  if (val.expire !== 0) {
    if (Date.now() > val.time + val.expire) {
      localStorage.removeItem(key);
      return null;
    }
  }

  return val.data;
}

export { SetLocalStorageItem, GetLocalStorageItem };
