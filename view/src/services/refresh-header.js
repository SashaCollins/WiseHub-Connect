import { secure } from './encryption.service';
import { cookies } from '@/cookies/cookies';

export default function refreshHeader() {
  let refresh = cookies.get('refresh');

  if (refresh) {
    return {
      'Authorization': 'Bearer ' + refresh,
    };
  } else {
    return {};
  }
}
