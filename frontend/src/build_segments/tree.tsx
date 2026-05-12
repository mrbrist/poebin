import { useState } from "react";
import type { buildData } from "../api/getBuild";

// This might need rewriting because the tree selection dropdown is using the title and not a uuid of some sort ¯\_(ツ)_/¯

function BuildTree({ build }: { build: buildData | undefined }) {
    const trees = build?.Data?.Tree?.Specs;

    const [selectedTreeTitle, setSelectedTreeTitle] = useState(trees[0].Title);
    const selectedTree = trees.find(
        (s: { Title: any }) => s.Title === selectedTreeTitle,
    );

    if (!selectedTree) return null;

    return (
        <div className="block text-left">
            {trees.length > 1 ? (
                <select
                    value={selectedTreeTitle}
                    onChange={(e) => setSelectedTreeTitle(e.target.value)}
                    className="mb-4 p-2 border rounded"
                >
                    {trees.map((set: any) => (
                        <option key={set.Title} value={set.Title}>
                            {set.Title}
                        </option>
                    ))}
                </select>
            ) : null}

            {/* Items */}
            <div>{selectedTree.Nodes}</div>
        </div>
    );
}

export default BuildTree;
