import 'dart:html';
import 'package:g0/g0.dart';

G0 go;

void main() {
  Api api = new FixtureApi();

  Element container = querySelector('#container');
  go = new G0(container, api);
}