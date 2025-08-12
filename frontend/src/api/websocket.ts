import { useEffect, useRef, useState } from "react";

type Message = {
  action: string;
  session_id: string;
  data: any;
}

function useWebSocket(url: string, onMessage: (msg: Message) => void) {
  const ws = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);
  const reconnectAttempts = useRef(0);
  const maxReconnectAttempts = 5;

  const connect = () => {
    ws.current = new WebSocket(url);

    ws.current.onopen = () => {
      console.log("WebSocket connected");
      setIsConnected(true);
      reconnectAttempts.current = 0;
    };

    ws.current.onclose = (e) => {
      console.log("WebSocket closed", e);
      setIsConnected(false);
      if (e.code !== 1000 && reconnectAttempts.current < maxReconnectAttempts) {
        const delay = Math.min(
          1000 * Math.pow(2, reconnectAttempts.current),
          30000
        );
        console.log("Reconnecting in", delay, "ms...");
        setTimeout(connect, delay);
        reconnectAttempts.current++;
      }
    };

    ws.current.onerror = (err) => {
      console.error("WebSocket error", err);
      ws.current?.close();
    };

    ws.current.onmessage = (e) => {
      try {
        const data = JSON.parse(e.data);
        onMessage(data);
      } catch (error) {
        console.error("Error parsing message", error);
      }
    };
  };

  useEffect(() => {
    connect();

    return () => {
      if (ws.current) {
        ws.current.close(1000, "Component unmounted");
      }
    };
  }, [url]);

  const sendMessage = (message: string) => {
    if (isConnected && ws.current) {
      try {
        ws.current.send(message);
      } catch (error) {
        console.error("Error sending message", error);
      }
    } else {
      console.warn("Cannot send message, WebSocket is not connected");
    }
  };

  return { sendMessage, isConnected };
}

export { useWebSocket };
