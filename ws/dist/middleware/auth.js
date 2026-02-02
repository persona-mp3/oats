function mockSharedDB() {
    const _masterUserName = "childish_gambino";
    const db = new Map();
    const masterUser = {
        authToken: "don't_be_scandalous",
        userName: "childish_gambino",
        password: "awaken_my_love",
        email: "donaldglover@spotify.com"
    };
    db.set(_masterUserName, masterUser);
    return db;
}
function masterCompadres() {
    const friends = { name: "jia_tian", status: false, lastSeen: "" };
    return new Map([["jia_tian", friends]]);
}
class Database {
    #users;
    constructor() {
        this.#users = mockSharedDB();
    }
    simulateConnection() {
        let counter = 0;
        let interval = setInterval(function () {
            counter += 1;
            if (counter === 2) {
                console.log("connected to the database");
                clearInterval(interval);
            }
            console.log("connecting to database...");
        }, 1300);
        console.log("connected to databse");
        return true;
    }
    findUser(userName, bearerToken) {
        let authStats = {
            status: false,
        };
        const userExists = this.#users.has(userName);
        if (!userExists) {
            console.error("user doesn't exist", userName, userExists);
            return authStats;
        }
        console.log("user exists, extracting data");
        const user = this.#users.get(userName);
        if (!user) {
            console.error("little patience for ts, but incase something like this happens...");
            throw new Error("little patience for ts, but incase something like this happens...");
        }
        console.log("extracted-token", bearerToken);
        if (user.authToken === bearerToken) {
            console.log("tokens match");
            authStats.info = masterCompadres();
            authStats.status = true;
            return authStats;
        }
        return authStats;
    }
}
const DB = new Database();
export function authClient(userName, bearerToken) {
    const isConnected = DB.simulateConnection();
    if (!isConnected)
        throw new Error("Yoo twin....Database could not be connected yo, what shall ye be done????😭✌️");
    return DB.findUser(userName, bearerToken);
}
