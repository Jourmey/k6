export let options = {
	stages: [
		{duration: '1s', target: '10'},
		{duration: '2s', target: '10'},
	],
};

export default function (data) {
	console.log("data:", data)
	// let res = http.get(data.url);
	// console.log("res:", res)

	data.response["data4"] = "lalalalalal4444"

	return data
};
