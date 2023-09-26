type WavesSpinnerProps = {
  color: "red" | "amber" | "indigo";
};

export function WavesSpinner(props: WavesSpinnerProps): JSX.Element {
  const colors: Record<WavesSpinnerProps["color"], string> = {
    indigo: "bg-indigo-500",
    amber: "bg-amber-500",
    red: "bg-red-500",
  };
  return (
    <div className="flex items-center justify-between relative w-2.5 h-4">
      <div className={`w-0.5 animate-spinner ${colors[props.color]}`} />
      <div
        className={`w-0.5 animate-spinner animation-delay-100 ${
          colors[props.color]
        }`}
      />
      <div
        className={`w-0.5 animate-spinner animation-delay-200 ${
          colors[props.color]
        }`}
      />
    </div>
  );
}
