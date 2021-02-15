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
        Ði {name} buvo pridëta {version} ir ankstesnëse versijose neveiks!
      </p>
    </Admonition>
  );
}
