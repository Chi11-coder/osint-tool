import { useState } from "react";
import { SearchForm } from "./components/SearchForm";
import { BiArrowBack } from "react-icons/bi";
import {
  SEARCH_TYPE,
  type SearchResponse,
  type SearchTypeValue,
} from "./types/common.types";
import type { FileReport } from "./types/file.types";
import { HostReportCompoent } from "./components/reports/host/HostReportComponent";
import type { HostReport } from "./types/host.types";
import { FileReportFlowProvider } from "./components/reports/file/FileReportComponent";

function App() {
  const [result, setResult] = useState<SearchResponse | null>(null);
  const [type, setType] = useState<number>(SEARCH_TYPE.HOST_SEARCH);
  const [mode, setMode] = useState<"search" | "summary">("search");
  const [errMsg, setErrMsg] = useState<string>("");

  const handleSearchResponse = (data: SearchResponse) => {
    setResult(data);
    setMode("summary");
  };

  const handleBackSearchPage = () => {
    setMode("search");
  };

  const handleSearchType = (searchType: SearchTypeValue) => {
    const validTypes: SearchTypeValue[] = [
      SEARCH_TYPE.HOST_SEARCH,
      SEARCH_TYPE.FILE_SEARCH,
    ];

    if (!validTypes.includes(searchType)) {
      setMode("search");
      setErrMsg(`unexpected search type: ${searchType}`);
      return;
    }
    setType(searchType);
    setErrMsg("");
  };

  const renderContent = () => {
    if (mode === "search") {
      return (
        <SearchForm
          onSearchReports={handleSearchResponse}
          onSearchType={handleSearchType}
          errMsg={errMsg}
        />
      );
    }

    if (type === SEARCH_TYPE.HOST_SEARCH) {
      return (
        <HostReportCompoent
          reports={result as HostReport}
          onBack={handleBackSearchPage}
        />
      );
    }

    if (type === SEARCH_TYPE.FILE_SEARCH) {
      return result ? (
        <FileReportFlowProvider
          reports={result as FileReport}
          onBack={handleBackSearchPage}
        />
      ) : (
        <></>
      );
    }
  };
  return <div>{renderContent()}</div>;
}

interface DashboardTitleProps {
  onBack: () => void;
}
export function DashboardTitle({ onBack }: DashboardTitleProps) {
  return (
    <button type="button" className="btn-search-top" onClick={onBack}>
      <BiArrowBack />
      <span>検索画面に戻る</span>
    </button>
  );
}

export default App;
