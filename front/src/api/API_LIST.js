const HOST = 'http://127.0.0.1:8181';

const LOGIN = HOST + '/api/admin/auth/login';
const ME = HOST + '/api/admin/auth/me';

const POSTS = HOST + '/api/admin/posts';

const CATEGORIES = HOST + '/api/admin/categories-all';
const CATEGORIES_LIST = HOST + '/api/admin/categories';

const PAGES = HOST + '/api/admin/pages';

const DASHBOARDS = HOST + '/api/admin/dashboards';



const Content = HOST + '/api/mindoc/content/10'
export default {
  LOGIN,
  ME,
  POSTS,
  CATEGORIES,
  CATEGORIES_LIST,
  PAGES,
  DASHBOARDS,
  Content
};
