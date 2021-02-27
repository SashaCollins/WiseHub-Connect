import VueCookies from 'vue-cookies'

VueCookies.config('7d', '', '' ,true, 'Lax');

function json() {
    return {
        to_json: function (value) {
            return JSON.stringify(value);
        },
        from_json: function (value) {
            return JSON.parse(value);
        }
    }
}

function cookies() {
    return {
        set: function(key, value) {
            VueCookies.set(key , value);
        },
        get: function(key) {
            return VueCookies.get(key);
        },
        remove: function(key) {
            VueCookies.remove(key);
        },
        exists: function(key) {
            return VueCookies.isKey(key);
        },
        all: function() {
            return VueCookies.keys();
        }
    };
}

export { cookies, json };