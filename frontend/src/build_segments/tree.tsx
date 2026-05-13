import { useEffect, useState } from "react";
import type { buildData } from "../api/getBuild";

function BuildTree({ build }: { build: buildData | undefined }) {
    const trees = build?.Data?.Tree?.Specs ?? [];

    const [selectedTreeTitle, setSelectedTreeTitle] = useState<string>("");

    // initialize once build loads
    useEffect(() => {
        if (trees.length > 0 && !selectedTreeTitle) {
            setSelectedTreeTitle(trees[0].Title);
        }
    }, [trees]);

    if (!build || trees.length === 0) return <div>Loading...</div>;

    const selectedTree = trees.find((s: any) => s.Title === selectedTreeTitle);

    if (!selectedTree) return null;

    return (
        <div className="block text-left">
            {trees.length > 1 && (
                <select
                    value={selectedTreeTitle}
                    onChange={(e) => setSelectedTreeTitle(e.target.value)}
                    className="mb-4 rounded border p-2"
                >
                    {trees.map((set: any) => (
                        <option key={set.Title} value={set.Title}>
                            {set.Title}
                        </option>
                    ))}
                </select>
            )}

            <div>{selectedTree.Nodes}</div>
        </div>
    );
}

export default BuildTree;
