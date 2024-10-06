export namespace api {
	
	export class Request {
	    Method: string;
	    URL: string;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Method = source["Method"];
	        this.URL = source["URL"];
	    }
	}
	export class Response {
	    ErrorMsg: string;
	    Body: number[];
	    Headers: {[key: string]: string[]};
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorMsg = source["ErrorMsg"];
	        this.Body = source["Body"];
	        this.Headers = source["Headers"];
	    }
	}

}

