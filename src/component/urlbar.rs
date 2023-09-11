use strum::IntoEnumIterator;
use tui_menu::{MenuItem, MenuState};

#[derive(Debug)]
pub struct UrlBar {
    pub title: String,
    pub text: String,
    pub cursor_position: usize,
    pub input_mode: InputMode,
    pub method: Method,
    pub method_menu: MenuState<Method>,
}

impl Default for UrlBar {
    fn default() -> Self {
        Self {
            title: String::from("URL"),
            text: String::from(""),
            cursor_position: 0,
            input_mode: InputMode::Normal,
            method: Method::Get,
            method_menu: MenuState::new(vec![MenuItem::group(
                Method::default().to_string(),
                Method::iter()
                    .map(|m| MenuItem::item(m.to_string(), m))
                    .collect(),
            )]),
        }
    }
}

#[derive(Debug, Default)]
pub enum InputMode {
    #[default]
    Normal,
    Insert,
}

#[derive(Debug, Default, Clone, strum::Display, strum::EnumIter)]
pub enum Method {
    #[default]
    Get,
    Post,
    Put,
    Patch,
    Delete,
    Head,
    Options,
}
