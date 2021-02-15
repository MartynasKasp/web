import Admonition from "../../Admonition";

export default function NoteLowercase({ name = "funkcija" }) {
  return (
    <Admonition type="warning">
      <p>Ði {name} prasideda maþàja raide.</p>
    </Admonition>
  );
}
