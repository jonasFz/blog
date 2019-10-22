class API {
	host: string;

	constructor(host: string){
		this.host = "http://" + host;
	}
	
	async hitEndpoint(endpoint: string, method: string): Promise<string>{
		return new Promise<string>((resolve, reject) => {

			let req = new XMLHttpRequest();
			console.log(this.host + endpoint);
			req.open(method, this.host+endpoint);
			req.onload = () =>{
				resolve(req.responseText);
			};
		
			req.send();
		});
	}

	async get(endpoint: string): Promise<string>{
		return await this.hitEndpoint(endpoint, "GET");
	}

	fullPath(endpoint: string): string{
		return this.host + endpoint;
	}
}

const api = new API("localhost:7991");

async function test(){
	let articles = (await api.get("/articles")).split("\n");
	
	let newStuff = "";

	for (let a of articles){
		if (a == ''){
			continue;
		}
		let parts = a.split(" ");
		let html = `<a href=${api.fullPath(parts[2])}>${parts[0]}</a>`;
		newStuff = newStuff + html + "\n";
	}
	let element = document.getElementById("content");
	if (element != null){
		element.innerHTML += newStuff;
	}
}

window.onload = test;
