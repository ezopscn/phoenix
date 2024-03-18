const MenuPermissionCheck = (path, menus) => {
  // 错误页面不需要验证
  let errUrlList = ['/error/403', '/error/404', '/error/500'];
  if (errUrlList.includes(path)) {
    return true;
  } else {
    for (const menu of menus) {
      if (menu.path === path) {
        return true;
      }
    }
  }
  return false;
};

export default MenuPermissionCheck;
