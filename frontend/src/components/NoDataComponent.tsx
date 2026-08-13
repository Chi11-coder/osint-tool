interface NoDataProps {
  name: string;
}
export function NoDataComponent({ name }: NoDataProps) {
  return (
    <div>
      <p>{name} のレポートデータの取得が正常に行えませんでした。</p>
    </div>
  );
}
