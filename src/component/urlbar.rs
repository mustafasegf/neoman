#[derive(Debug, Default)]
pub struct UrlBar {
    pub title: String,
    pub text: String,
    pub cursor_position: usize,
    pub input_mode: InputMode,
    pub method: Method,
}

#[derive(Debug, Default)]
pub enum InputMode {
    #[default]
    Normal,
    Insert,
}

#[derive(Debug, Default, strum::Display)]
pub enum Method {
    #[default]
    Get,
    Post,
    Put,
    Patch,
    Delete,
    Head,
    Options,
    Other(String),
}

