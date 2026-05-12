import { type buildData } from "../api/getBuild";

function BuildNotes({ build }: { build: buildData | undefined }) {
    if (build?.Data.Notes.trim() == "") {
        return <div></div>;
    } else {
        return (
            <div className="col-span-9 rounded-2xl border border-slate-400 p-4">
                <h2 className="text-lg font-semibold mb-4">Notes</h2>
                <span className="whitespace-pre-line block text-left">
                    {build?.Data.Notes}
                </span>
            </div>
        );
    }
}

export default BuildNotes;
