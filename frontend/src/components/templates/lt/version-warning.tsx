import Admonition from "../../Admonition";

export default function WarningVersion({
  version,
  name = "funkcija",
}: {
  version: string;
  name: string;
}) {
  return (
    <Admonition type="warning">
      <p>
        �i {name} buvo prid�ta {version} ir ankstesn�se versijose neveiks!
      </p>
    </Admonition>
  );
}
