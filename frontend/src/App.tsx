import { useEffect, useState } from "react";
import { newBuild } from "./api/newBuild";
import { useNavigate } from "react-router-dom";
import { getRecent, type buildDataArr, type buildData } from "./api/getRecent";
import { timeAgo } from "./helpers/timeAgo";

async function handlePaste(paste: string, nav: any) {
    if (!paste) {
        return;
    }
    const id = await newBuild(paste);
    if (!id) {
        return;
    }
    nav("/" + id);
}

async function handleRecent(set: any) {
    const recent = await getRecent();
    if (!recent) {
        return;
    }
    set(recent);
}

function App() {
    const [paste, setPaste] = useState("");
    const [recent, setRecent] = useState<buildDataArr>();
    const navigate = useNavigate();

    useEffect(() => {
        const init = async () => {
            await handleRecent(setRecent);
        };
        init();
    }, []);

    return (
        <div className="h-screen flex flex-col">
            <div className="flex justify-center p-4">
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
                        onClick={() => handlePaste(paste, navigate)}
                    >
                        Create
                    </button>
                </div>
            </div>

            <div className="flex-1 flex justify-center">
                <div className="w-full max-w-2xl">
                    <h2 className="pt-5 text-center text-lg font-semibold text-slate-200">
                        Recent Builds
                    </h2>

                    <div className="mt-4 space-y-2">
                        {recent?.map((r: buildData) => (
                            <a
                                key={r.ID}
                                href={"/" + r.ID}
                                className="flex items-center gap-3 rounded-lg bg-slate-800/70 px-4 py-3 transition hover:bg-slate-700/70"
                            >
                                <img
                                    src="https://placehold.co/100"
                                    alt="Class avatar"
                                    className="h-10 w-10 rounded-full object-cover"
                                />

                                <div className="flex flex-col text-left">
                                    <span className="text-sm font-medium text-slate-100">
                                        Level {r.Level} {r.Class}
                                    </span>
                                    <span className="text-xs text-slate-400">
                                        {timeAgo(r.DateAdded)}
                                    </span>
                                </div>
                            </a>
                        ))}
                    </div>
                </div>
            </div>

            <footer className="bg-neutral-primary-soft flex justify-center items-center p-4">
                <span className="text-sm text-slate-500 text-center">
                    poeb.in isn't affiliated with or endorsed by Grinding Gear
                    Games in any way
                </span>
            </footer>
        </div>
    );
}

export default App;
