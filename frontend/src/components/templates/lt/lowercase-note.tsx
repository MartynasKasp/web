import Admonition from "../../Admonition";

export default function NoteLowercase({ name = "funkcija" }) {
  return (
    <Admonition type="warning">
      <p>�i {name} prasideda ma��ja raide.</p>
    </Admonition>
  );
}
