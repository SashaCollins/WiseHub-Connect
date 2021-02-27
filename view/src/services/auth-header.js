import { secure } from './encryption.service';

export default function authHeader() {
  let token = secure.get('token');

  if (token) {
    return {
      'Authorization': 'Bearer ' + token
    };
  } else {
    return {};
  }
}
