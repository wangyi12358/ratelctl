export interface BaseProps {
  className?: string;
  style?: React.CSSProperties;
}

export interface BaseFormProps<V = string> extends BaseProps {
  value?: V;
  defaultValue?: V;
  onChange?: (value: V) => void;
}