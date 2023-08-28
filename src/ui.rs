use ratatui::{
    prelude::*,
    widgets::{Block, BorderType, Borders, Paragraph, Tabs, Wrap},
};
use strum::IntoEnumIterator;
use tui_tree_widget::Tree;

use crate::{app::{App, Selected}, component::requestbar::RequestMenu};

pub const HIGHLIGHT_STYLE: Style = Style::new()
    .fg(Color::LightBlue)
    .add_modifier(Modifier::BOLD)
    .bg(Color::DarkGray);

pub const SELECTED_STYLE: Style = Style::new().fg(Color::LightGreen);

pub const DEFAULT_STYLE: Style = Style::new().fg(Color::White);

/// Renders the user interface widgets.
pub fn render<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>) {
    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .constraints([Constraint::Length(app.sidebar_size()), Constraint::Min(0)].as_ref())
        .split(frame.size());

    sidebar(app, frame, chunks[0]);
    mainbar(app, frame, chunks[1]);
}

pub fn sidebar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let (style, highlight_style) = match app.selected == Selected::Sidebar {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let block = Block::default()
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .style(style);

    let items = Tree::new(app.sidebar.tree.items.clone())
        .block(block)
        .highlight_style(highlight_style)
        .style(DEFAULT_STYLE);

    frame.render_stateful_widget(items, area, &mut app.sidebar.tree.state);
}

pub fn mainbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints(
            [
                Constraint::Min(3),
                Constraint::Min(3),
                Constraint::Ratio(1, 2),
                Constraint::Ratio(1, 2),
            ]
            .as_ref(),
        )
        .split(area);

    tabs(app, frame, chunks[0]);
    urlbar(app, frame, chunks[1]);
    requestbar(app, frame, chunks[2]);
    responsebar(app, frame, chunks[3]);
}

pub fn tabs<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let (style, highlight_style) = match app.selected == Selected::Tabs {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let titles = app
        .tabs
        .tabs
        .iter()
        .map(|item| Line::from(item.to_string()))
        .collect();

    let tabs = Tabs::new(titles)
        .block(Block::default().borders(Borders::ALL).title("Tabs"))
        .select(app.tabs.selected)
        .style(style)
        .highlight_style(highlight_style);

    frame.render_widget(tabs, area);
}

pub fn urlbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let (method_style, _method_highlight_style) = match app.selected == Selected::MethodBar {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let (url_style, _url_highlight_style) = match app.selected == Selected::Urlbar {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .constraints([Constraint::Min(10), Constraint::Min(0)].as_ref())
        .split(area);

    let block = Block::default()
        .title(format!("URL: {}", app.urlbar.title))
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .style(url_style);

    let text = Paragraph::new(app.urlbar.text.clone())
        .block(block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    let method_block = Block::default()
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .style(method_style);

    let method = Paragraph::new(app.urlbar.method.to_string())
        .block(method_block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    frame.render_widget(method, chunks[0]);
    frame.render_widget(text, chunks[1]);
}

pub fn requestbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let (style, highlight_style) = match app.selected == Selected::Requestbar {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints([Constraint::Min(1), Constraint::Min(0)].as_ref())
        .split(area);

    let titles = RequestMenu::iter()
        .map(|item| Line::from(item.to_string()))
        .collect();

    let tabs = Tabs::new(titles)
        // .block(Block::default())
        .select(0)
        .style(style)
        .highlight_style(highlight_style);

    frame.render_widget(tabs, chunks[0]);

    let block = Block::default()
        .title("Request")
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);

    let text = Paragraph::new(app.requestbar.body.clone())
        .block(block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    frame.render_widget(text, chunks[1]);
}

pub fn responsebar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let (style, _highlight_style) = match app.selected == Selected::Responsebar {
        true => (SELECTED_STYLE, HIGHLIGHT_STYLE),
        false => (DEFAULT_STYLE, DEFAULT_STYLE),
    };

    let block = Block::default()
        .title("Response")
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .style(style);

    let text = Paragraph::new(app.responsebar.body.clone())
        .block(block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    frame.render_widget(text, area);
}
