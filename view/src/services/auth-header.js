import { secure } from './encryption.service';

export default function authHeader() {
  let token = secure.get('user');

  if (token) {
    return {
      'Authorization': 'Bearer ' + token
    };
  } else {
    return {};
  }
}
