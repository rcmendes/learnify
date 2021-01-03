#[derive(Debug)]
pub enum AppError {
    HttpRequestError(String),
    IOError(String),
}


impl From<reqwest::Error> for AppError {
    fn from(err: reqwest::Error) -> Self {
        Self::HttpRequestError(err.to_string())
    }
}

impl From<std::io::Error> for AppError {
    fn from(err: std::io::Error) -> Self {
        Self::IOError(err.to_string())
    }
}

impl std::fmt::Display for AppError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let message = match self {
            Self::HttpRequestError(msg) => {
                format!("An error occurred when handling request. Detail: {}", msg)
            }

            Self::IOError(msg) => {
                format!("An error occurred in IO processing. Detail: {}", msg)
            } //
        };

        write!(f, "{}", message)
    }
}

impl std::error::Error for AppError {}
