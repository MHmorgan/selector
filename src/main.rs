extern crate anyhow;
#[macro_use]
extern crate clap;
extern crate cursive;

use anyhow::Result;
use cursive::align::HAlign;
use cursive::theme::BaseColor::Black;
use cursive::theme::Color::Dark;
use cursive::theme::PaletteColor::Background;
use cursive::theme::{BorderStyle, Palette, Theme};
use cursive::traits::Resizable;
use cursive::views::{Dialog, SelectView};

fn main() -> Result<()> {
    let m = clap_app!(myapp =>
        (version: "1.0")
        (author: "Magnus Aa. Hirth <magnus.hirth@gmail.com")
        (about: "Select between multiple arguments")
        (@arg VALS: ... +required "Values to select between")
    )
    .get_matches();

    let vals: Vec<&str> = m.values_of("VALS").unwrap().collect();

    let mut sel = SelectView::new().h_align(HAlign::Center);
    sel.add_all_str(vals);
    sel.sort();
    sel.set_on_submit(|s, v: &String| {
        s.quit();
        print!("{}", v)
    });

    let mut palette = Palette::default();
    palette[Background] = Dark(Black);
    let theme = Theme {
        shadow: false,
        borders: BorderStyle::None,
        palette: palette,
    };

    let mut siv = cursive::default();
    siv.set_theme(theme);
    siv.add_layer(Dialog::around(sel).min_width(20));
    siv.run();

    Ok(())
}
