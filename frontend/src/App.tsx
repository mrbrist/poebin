import { useState } from "react";

function handlePaste(paste: string) {
    if (!paste) {
        return;
    }
    console.log(paste);
}

function App() {
    const [paste, setPaste] = useState("");
    return (
        <div className="min-h-screen flex justify-center p-4">
            <div className="w-full max-w-2xl rounded-2xl p-6">
                <h1 className="text-2xl font-semibold mb-4 text-slate-100">
                    poeb.in
                </h1>

                <textarea
                    value={paste}
                    onChange={(e) => setPaste(e.target.value)}
                    placeholder="Paste Path of Building code here..."
                    className="w-full h-64 p-4 bg-slate-800 text-slate-100 placeholder-slate-500 rounded-xl resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"
                />

                <button
                    className="mt-2 px-5 py-2 rounded-xl bg-blue-600 text-white font-medium hover:bg-blue-500 active:bg-blue-700 transition w-10/12"
                    onClick={() => handlePaste(paste)}
                >
                    Create
                </button>
            </div>
        </div>
    );
}

export default App;
