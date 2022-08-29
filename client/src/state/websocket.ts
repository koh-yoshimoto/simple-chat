import { atom, selector } from "recoil";
import * as Websocket from "websocket";

const connect = (): Promise<Websocket.w3cwebsocket> => {
    return new Promise((resolve, reject) => {
        const port = import.meta.env.VITE_WS_PORT;
        const url = "ws://localhost:" + port + "/ws";
        const socket = new Websocket.w3cwebsocket(url);

        socket.onopen = () => {
            console.log("connected", port);
            resolve(socket);
        };

        socket.onclose = () => {
            console.log("reconnecting...");
            connect();
        };

        socket.onerror = (error) => {
            console.log("connection error:", error);
            reject(socket);
        };
    });
};

const connectWebsocketSelector = selector({
    key: "connectWebsocket",
    get: async ():Promise<Websocket.w3cwebsocket> => {
        return await connect();
    },
});

export const websocketAtom = atom<Websocket.w3cwebsocket>({
    key: "websocket",
    default: connectWebsocketSelector,
});
