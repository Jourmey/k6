import http from 'k6/http';

export default function () {
	return http.get('https://test-api.k6.io/');
};
