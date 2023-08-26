use std::error;

use tui_tree_widget::TreeItem;

use crate::items::{Item, StatefulTree};

/// Application result type.
pub type AppResult<T> = std::result::Result<T, Box<dyn error::Error>>;

/// Application.
#[derive(Debug)]
pub struct App {
    pub running: bool,
    pub selected: Selected,
    pub sidebar: SideBar,
    pub settings: Settings,
    pub tabs: TabBar,
    pub urlbar: UrlBar,
    pub requestbar: RequestBar,
    pub responsebar: ResponseBar,
}

#[derive(Debug, Default, strum::Display, strum::EnumIter, PartialEq)]
pub enum Selected {
    #[default]
    Sidebar,
    Tabs,
    MethodBar,
    Urlbar,
    Requestbar,
    Responsebar,
}

#[derive(Debug)]
pub struct SideBar {
    pub size: u16,
    pub selected: usize,
    pub tree: StatefulTree<'static>,
}

#[derive(Debug, Default)]
pub struct TabBar {
    pub selected: usize,
    pub tabs: Vec<Item>,
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

#[derive(Debug, Default)]
pub struct UrlBar {
    pub title: String,
    pub text: String,
    pub cursor_position: usize,
    pub input_mode: InputMode,
    pub method: Method,
}

#[derive(Debug, Default, strum::Display, strum::EnumIter)]
pub enum RequestMenu {
    #[default]
    Params,
    Authentication,
    Headers,
    Body,
}

#[derive(Debug, Default)]
pub struct RequestBar {
    pub body: String,
    pub request_menu: RequestMenu,
}

#[derive(Debug, Default)]
pub struct ResponseBar {
    pub body: String,
}

#[derive(Debug)]
pub struct Settings {
    pub show_sidebar: bool,
    pub show_help: bool,
}

impl Default for Settings {
    fn default() -> Self {
        Self {
            show_sidebar: true,
            show_help: false,
        }
    }
}

impl Default for App {
    fn default() -> Self {
        let mut tree = StatefulTree::with_items(vec![
            TreeItem::new_leaf(Item::new("a")),
            TreeItem::new(
                Item::new("b"),
                vec![
                    TreeItem::new_leaf(Item::new("c")),
                    TreeItem::new(
                        Item::new("d"),
                        vec![
                            TreeItem::new_leaf(Item::new("e")),
                            TreeItem::new_leaf(Item::new("f")),
                        ],
                    ),
                    TreeItem::new_leaf(Item::new("g")),
                ],
            ),
            TreeItem::new_leaf(Item::new("d")),
        ]);
        tree.first();

        let tabs = tree.items.iter().map(|i| i.inner().clone()).collect();

        Self {
            running: true,
            selected: Selected::Sidebar,
            sidebar: SideBar {
                size: 25,
                selected: 0,
                tree,
            },
            settings: Settings {
                show_sidebar: true,
                show_help: false,
            },
            tabs: TabBar { selected: 0, tabs },
            urlbar: UrlBar {
                title: String::from("localhost:8080"),
                text: String::from("localhost:8080"),
                cursor_position: 0,
                input_mode: InputMode::Normal,
                method: Method::Get,
            },
            requestbar: RequestBar {
                body: String::new(),
                request_menu: RequestMenu::Params,
            },
            responsebar: ResponseBar {
                body: String::new(),
            },
        }
    }
}

impl App {
    /// Constructs a new instance of [`App`].
    pub fn new() -> Self {
        Self::default()
    }

    /// Handles the tick event of the terminal.
    pub fn tick(&self) {}

    /// Set running to false to quit the application.
    pub fn quit(&mut self) {
        self.running = false;
    }

    pub fn toggle_sidebar(&mut self) {
        self.settings.show_sidebar = !self.settings.show_sidebar;
    }

    pub fn sidebar_size(&self) -> u16 {
        match self.settings.show_sidebar {
            true => self.sidebar.size,
            false => 0,
        }
    }
}
