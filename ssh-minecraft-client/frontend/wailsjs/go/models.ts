export namespace main {
	
	export class Request {
	    status: boolean;
	    meta: string;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.meta = source["meta"];
	    }
	}

}

