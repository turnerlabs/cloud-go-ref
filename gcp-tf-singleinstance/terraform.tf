// Configure the Google Cloud provider
provider "google" {
  credentials = "${file("/Users/mickey/Dropbox/my_docs/software/GCP/My First Project-42231c53a0e2.json")}"
  project     = "axiomatic-array-130823"
  region      = "us-central1"
}


resource "google_compute_firewall" "default" {
  name    = "eighty"
  network = "default"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["80", "8080"]
  }

  source_ranges = ["0.0.0.0/0"]
}


resource "google_compute_instance" "default" {
  name         = "cloud-go-ref"
  description = "test server for cloud-go-ref"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"

  tags = ["cloud-go-ref"]

  disk {
    image = "ubuntu-1404-trusty-v20160509a"
  }

  // Local SSD disk
  disk {
    type    = "local-ssd"
    scratch = true
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral IP
    }
  }

  metadata {
    billing = "TBD"
    environment = "dev"
    creator = "mickeyyawn@gmail.com"
    customer = "MSS-CA-DEVOPS"
    team = "MSS-CA-DEVOPS"
    product = "foo"
  }


  metadata_startup_script = "${file("gcp_startup_script.txt")}"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}
