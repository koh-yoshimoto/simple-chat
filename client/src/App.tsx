import { MessageInput } from "./components/MessageInput";
import { MessageList } from "./components/MessageList";

function App() {
  return (
    <div>
        <h1>simple chat</h1>
        <MessageInput />
        <MessageList />
    </div>
  );
};

export default App
