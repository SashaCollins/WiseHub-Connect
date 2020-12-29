### 1.2 Versionsverwaltung


Keine verantwortungsvolle Person würde heutzutage ein Projekt starten ohne eine backup Strategie. Denn Daten sind flüchtig und können leicht verloren gehen, entweder druch eine fehlerhafte Änderung oder einen fatalen Festplattenabsturz. Deshalb sollte jede Arbeit am Projekt archiviert werden. Für Software Projekte ist die typische backup Strategie ein Versionskontrollsystem, sprich das Verfolgen und Verwalten von Überarbeitungen und Änderungen.

https://books.google.de/books?hl=de&lr=&id=aM7-Oxo3qdQC&oi=fnd&pg=PR3&dq=software+version+control&ots=3atlIJYduh&sig=1M0dWFsJ-5X2rqS9Q5pygbnqFiY#v=onepage&q=software%20version%20control&f=false
29.12.2020 01:53

Das System der Versionsverwaltung ist in der Lage, Arbeitskopien (organisiert in einem Verzeichnisbaum) mit den Daten aus dem Repository (engl. Aufbewahrungsort) zu synchronisieren. Das ist nötig, da es nur so möglich ist, mit den im Repository abgelegten Dateien zu arbeiten. Wird eine Version aus dem Repository in die Arbeitskopie übertragen, wird von einem Checkout, von Aus-Checken oder von Aktualisieren gesprochen. Die umgekehrte Übertragung von der Arbeitskopie in das Repository hingegen wird als Check-in, als Einchecken oder als Commit bezeichnet. Nun können die im Repository abgelegten Dateien mit diversen Programmen bearbeitet werden. Hierbei kann es sich um kommandozeilenorientierte Programme oder um solche mit grafischer Benutzeroberfläche handeln. In einigen Fällen kann das Programm aber natürlich auch in Form eines Plugins für die integrierte Softwareentwicklungsumgebung eingesetzt werden. In der Regel haben Nutzer die Wahl zwischen mehreren dieser Zugriffsmöglichkeiten.

Die Versionsverwaltung macht es möglich, Änderungen an einer Datei oder an einer Gruppe von Dateien zu speichern. Allerdings wird hier nicht etwa nur die Version nach der letzten Änderung gespeichert. Zunächst wäre da die Ausgangsversion einer Datei, die gesichert wird. Wird nun eine Änderung vollzogen, wird auch diese Änderung gespeichert, allerdings bleibt gleichzeitig auch die Ausgangsversion bestehen. So ist es für den Nutzer möglich, immer auf eine vorherige Version zurückzugreifen, wenn es Probleme mit der aktuellen Version gibt. Gleichzeitig ist die Versionsverwaltung aber auch essenziell, um überhaupt mit den Dateien arbeiten zu können. Das System sorgt nämlich dafür, dass der aktuelle (sowie der ältere) Stand eines Projekts, das in Form eines Verzeichnisbaums gespeichert ist, mit dem Repository synchronisiert wird. Aufgrund der verschiedenen Versionsverwaltungssysteme gibt es darüber hinaus für jeden Anwendungsbereich ein passendes System.

Derzeit gibt es insgesamt drei Arten der Versionsverwaltung. Das System ist hier entweder:

- lokal  
  Lokale Versionsverwaltungssysteme wurden mit Werkzeugen wie RCS und SCCS umgesetzt. Sie funktionieren nur auf einem Computer. Hinzu kommt, dass die lokale Versionsverwaltung meist nur eine einzige Datei versioniert. Anwendung findet die lokale Versionsverwaltung vor allem in Büroanwendungen. Hier werden die Versionen eines Dokuments dann direkt in der Datei des Dokuments gespeichert. In technischen Zeichnungen hingegen erfolgt die Versionsverwaltung hingegen etwa durch einen Änderungsindex.

- zentral  
Die zentrale Versionsverwaltung ist als Client-Server-System aufgebaut und erlaubt so den netwerkweiten Zugriff auf ein Repositry. Mithilfe einer Rechteverwaltung kann dafür gesorgt werden, dass nur berechtigte Personen eine neue Version in das Archiv legen können. Populär wurde dieses Konzept durch das Open-Source-Projekt Concurrent Versions System (CVS). Neu implementiert hingegen wurde es mit Subversion (SVN) und wird inzwischen von vielen kommerziellen Anbietern eingesetzt.

- verteilt  
Die verteilte Versionsverwaltung hingegen verwendet kein zentrales Repository. So verfügt jeder, der an einem Projekt arbeitet, ein eigenes Repository, das er mit jedem x-beliebigen anderem Repository abgleichen kann. So ist die Versionsgeschichte genauso verteilt wie in den zentralen Verwaltungssystemen. Der Vorteil ist allerdings, dass Änderungen auch lokal erfolgen können - ganz ohne den Aufbau einer Server-Verbindung.
Ihre Gemeinsamkeit: alle drei Generationen speichert üblicherweise nur die Unterschiede zwischen zwei Versionen. So lässt sich Speicherplatz sparen.

Sollten mehrere Benutzer dieselbe Version einer Datei verändern, kann es so zu keinem Konflikt zwischen den einzelnen Versionen kommen. Sich widersprechende Versionen existieren anfangs parallel nebeneinander und können weiterhin verändert werden. Später ist es ebenso möglich, die Versionen in eine neue Version zusammenzuführen. Anstelle einer Kette von Versionen entsteht so eine Polyhierarchie (gerichteter azyklischer Graph). Diese Möglichkeit macht sich die Softwareentwicklung zum Vorteil. So werden zunächst einzelne Features oder Feature-Gruppen in separaten Versionen entwickelt. Anschließend werden sie alle von Personen mit einer Integrator-Rolle überprüft und später zusammengeführt.

https://www.it-talents.de/blog/it-talents/was-ist-eine-versionsverwaltung 29.12.2020 02:04

Git beispielsweise ist ein solches VCS.
Für die anfertigung dieser Bachelorarbeit wurde angenommen das als VCS GitHub verwendet wird.
GitHub bietet einen Cloud-basierten Git Repository Hosting Service. Im Wesentlichen macht es Einzelpersonen und Teams viel einfacher, Git für Versionskontrolle und Zusammenarbeit zu nutzen.
