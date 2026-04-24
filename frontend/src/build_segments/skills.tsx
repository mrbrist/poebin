import type { buildData } from "../api/getBuild";

function BuildSkills({ build }: { build: buildData | undefined }) {
    return <div>{JSON.stringify(build?.Data.Skills.SkillSets[0])}</div>;
}

export default BuildSkills;
